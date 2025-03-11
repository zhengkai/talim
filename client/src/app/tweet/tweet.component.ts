import { Component, input, effect } from '@angular/core';
import { api } from '../common/api';
import { pb } from '../../pb';
import Long from 'long';
import { DatePipe } from '@angular/common';
import { TweetItemComponent } from './item.component';

const LongZero = Long.fromNumber(0);

@Component({
	selector: 'app-tweet',
	imports: [
		DatePipe,
		TweetItemComponent,
	],
	templateUrl: './tweet.component.html',
})
export class TweetComponent {

	loadDone = false;

	userMap: { [key: string]: pb.User } = {};

	cover: pb.User | null = null;

	uid = input.required<string>();

	api = api;

	loadCnt = 0;

	tweet: pb.TweetRow[] = [];

	constructor() {
		effect(async () => {
			this.load(this.uid());
		});
	}

	async load(uid: string, tid = LongZero) {
		this.loadCnt++;
		if (this.loadCnt > 4) {
			return;
		}
		const o = await api.tweet(Long.fromString(uid), tid);
		this.loadDone = true;
		if (!o?.tweet?.length) {
			return;
		}
		for (const r of o?.user || []) {
			const u = pb.User.fromObject(r);
			const suid = u.uid.toString();
			if (suid === uid) {
				this.cover = u;
				console.log(u);
			}
			this.userMap[suid] = u;
		}
		const len = o?.tweet?.length || 0;
		if (len) {
			for (const r of o.tweet || []) {
				const t = pb.TweetRow.fromObject(r);
				this.tweet.push(t);
				const doc = new DOMParser().parseFromString(t.text, "text/html");
				t.text = doc.documentElement.textContent || '';
			}
			if (len === 5000) {
				const last = o.tweet[len - 1]?.tid;
				if (last) {
					this.load(uid, last);
					console.log(o.tweet[0].tid?.toString());
					console.log(last.toString())
				}
			}
		}
	}
}
