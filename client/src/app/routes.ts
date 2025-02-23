import { Routes } from '@angular/router';

import { IndexComponent } from './view/index.component';
import { TweetComponent } from './tweet/tweet.component';
import { TweetRecentComponent } from './tweet/recent.component';

export const routes: Routes = [
  { path: 'recent', component: TweetRecentComponent },
  { path: 'tweet/:uid', component: TweetComponent },
  { path: 'home', component: IndexComponent },
  { path: '**', redirectTo: '/home' },
];
