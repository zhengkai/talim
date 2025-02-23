import { Component } from '@angular/core';
import { api } from '../common/api';

@Component({
  selector: 'app-index',
  imports: [],
  templateUrl: './index.component.html',
})
export class IndexComponent {

  api = api;

}
