// ==UserScript==
// @name         Twitter Data Copier
// @namespace    https://talim.9farm.com/
// @version      1.0
// @description  搬运 twtter 数据
// @author       Soulogic
// @match        https://x.com/*
// @connect      talim.9farm.com
// @grant        GM_xmlhttpRequest
// ==/UserScript==

//
//
//
//
// 警告：本脚本收集所有 twitter(x.com) 的 api 数据
//
// 我的本意是收集和备份 tweet 数据，但我没时间研究每一个接口
//
// 因此可能会收集到敏感数据
//
// 我强烈建议自行搭建 server 端，
// 并将本脚本里的 talim.9farm.com 替换为你自己的域名
//
//
//
//

(() => {
	'use strict';

	// 如果不想跟其他人共享数据，可以自行生成个 uuid
	const uuid = '54c2d184-77eb-4333-abfa-91b002e76827';

	const send = async function(data) {
		if (!(data?.length > 1000)) {
			return;
		}
		GM_xmlhttpRequest({
			url: `https://talim.9farm.com/api/upload?uuid=${uuid}`,
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			data,
		});
	};

	const originalOpen = XMLHttpRequest.prototype.open;
	XMLHttpRequest.prototype.open = function(method, url, ...rest) {
		this.addEventListener('readystatechange', function() {
			if (!url.includes('notifications/all.json') && !url.includes('api/graphql')) {
				return
			}
			if (this.readyState === 4 && this.status === 200) {
				send(this.responseText);
			}
		});
		return originalOpen.apply(this, [method, url, ...rest]);
	};
})();
