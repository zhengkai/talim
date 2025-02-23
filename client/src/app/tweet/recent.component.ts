import { Component, input, effect } from '@angular/core';
import { api } from '../common/api';
import { OnInit } from '@angular/core';

@Component({
  selector: 'app-tweet-recent',
  imports: [],
  templateUrl: './tweet-recent.component.html',
})
export class TweetRecentComponent implements OnInit {

  loadDone = false;

  uid = input.required<string>();

  api = api;

  constructor() {
    effect(() => {
      let h = this.uid()
      console.log(h);
      // this.load();
    });
  }

  load(uid: string) {
    console.log('load', uid)
  }

  ngOnInit() {
    console.log('TweetRecentComponent initialized');
  }
}
