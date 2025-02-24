// ==UserScript==
// @name         Twitter Data Loop Copier
// @namespace    https://talim.9farm.com/
// @version      1.0
// @description  按日期倒刷 Twitter 数据，用于复制历史数据
// @author       Soulogic
// @match        https://x.com/*
// ==/UserScript==

(() => {
	'use strict';

	const user = 'soulogic';
	const stopDate = '2007-12-13';

	// location.href = '/search?q=' + encodeURIComponent('until:2022-10-01 from:soulogic') + '&f=live'

	let targetDate = ''; // targetDate 指看到这天的就翻页

	const getDateList = () => {
		const tli = [];
		document.querySelectorAll('time').forEach(o => {
			if (o.parentElement.tagName !== 'A') {
				return;
			}
			const s = o.getAttribute('datetime');
			tli.push(s.substring(0, 10));
		});
		tli.sort();
		return tli;
	}

	const refreshBtn = () => {
		for (const o of document.querySelectorAll('button span')) {
			if (o.innerText === '重试') {
				o.closest('button').click();
				return true;
			}
		};
		return false;
	}

	const getDate = () => {

		const tli = getDateList();
		console.log('开始扫描日期', tli)

		if (tli.length < 3) { // 可能页面没加载完
			console.log('页面不完全，等待重新扫描(2s)')
			setTimeout(() => {
				getDate();
			}, 2000);
			return;
		}

		if (targetDate === '') { // 最大日期总是当前日期
			const date = new Date(tli.pop());
			date.setDate(date.getDate() - 1);
			targetDate = date.toISOString().split('T')[0];
		}

		const minDate = tli.shift();
		console.log(`翻页日期 ${targetDate}，最小日期 ${minDate}`)
		if (minDate <= targetDate) {
			console.log(`准备翻页`)
			setTimeout(() => {
				nextPage(minDate);
			}, 6000);
			return;
		} else {
			console.log(`继续加载`)
			refreshBtn();
			window.scrollTo(0, document.body.scrollHeight);
			setTimeout(() => {
				getDate();
			}, 5000);
		}
	};

	const nextPage = (s) => {
		if (s <= stopDate) {
			console.log(`完成扫描，停止工作`)
			return;
		}
		const query = `until:${s} from:${user}`
		console.log(`刷新页面 ${query}`)
		window.location.href = '/search?q=' + encodeURIComponent(query) + '&f=live'
	}

	// 进页面 10 秒后开始扫描
	console.log('插件已加载，10 秒后开始扫描')
	setTimeout(() => {
		getDate();
	}, 50000);
})();
