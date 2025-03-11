import { Component } from '@angular/core';
import { api } from '../common/api';
import { pb } from '../../pb';
import { TweetItemComponent } from './item.component';

@Component({
	selector: 'app-tweet-recent',
	imports: [
		TweetItemComponent,
	],
	templateUrl: './recent.component.html',
})
export class TweetRecentComponent {

	loadDone = false;

	tweet: pb.TweetRow[] = [];

	api = api;

	constructor() {
		this.load();
	}

	async load() {
		const o = await api.recent();
		this.loadDone = true;
		if (!o?.tweet?.length) {
			return;
		}

		for (const r of o.tweet || []) {
			const t = pb.TweetRow.fromObject(r);
			this.tweet.push(t);
			const doc = new DOMParser().parseFromString(t.text, "text/html");
			t.text = doc.documentElement.textContent || '';
		}
	}
}
