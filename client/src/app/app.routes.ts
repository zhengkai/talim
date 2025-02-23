import { Routes } from '@angular/router';

import { IndexComponent } from './view/index.component';

export const routes: Routes = [
  { path: 'home', component: IndexComponent },
  { path: '**', redirectTo: '/home' },
];
