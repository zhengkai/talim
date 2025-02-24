import { Component, input, effect } from '@angular/core';
import { api } from '../common/api';
import { pb } from '../../pb';
import Long from 'long';
import { DatePipe } from '@angular/common';

@Component({
	selector: 'app-tweet',
	imports: [
		DatePipe,
	],
	templateUrl: './tweet.component.html',
})
export class TweetComponent {

	loadDone = false;

	userMap: { [key: string]: pb.User } = {};

	cover: pb.User | null = null;

	uid = input.required<string>();

	api = api;

	tweet: pb.TweetRow[] = [];

	constructor() {
		effect(async () => {
			this.load(this.uid());
		});
	}

	async load(uid: string) {
		const o = await api.tweet(Long.fromString(uid), Long.fromNumber(0));
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
		for (const r of o?.tweet || []) {
			const t = pb.TweetRow.fromObject(r);
			this.tweet.push(t);
			const doc = new DOMParser().parseFromString(t.text, "text/html");
			t.text = doc.documentElement.textContent || '';
		}
	}
}
