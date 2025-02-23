import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { pb } from '../../pb';
import Long from 'long';
import { api } from '../common/api';
import protobuf from 'protobufjs';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet],
  templateUrl: './bootstrap.component.html',
})
export class BootstrapComponent {

  title = 'Talim';

  loadDone = false;
  loadError = false;

  constructor(
  ) {
    protobuf.util.Long = Long;
    protobuf.configure();
    this.getUser()
  }

  async getUser(): Promise<void> {
    const li = await api.index();
    this.loadDone = true;
    if (!li.length) {
      this.loadError = true;
    }
  }
}
