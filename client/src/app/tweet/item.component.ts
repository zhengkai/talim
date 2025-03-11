import { Component, Input } from '@angular/core';
import { DatePipe } from '@angular/common';
import { pb } from '../../pb';

@Component({
	selector: 'app-tweet-item',
	imports: [
		DatePipe,
	],
	templateUrl: './item.component.html',
})
export class TweetItemComponent {

	@Input() t: pb.TweetRow | null = null;

	constructor() {
	}
}
