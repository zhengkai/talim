import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { pb } from '../../pb/pb';
import Long from 'long';
import protobuf from 'protobufjs';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet],
  templateUrl: './bootstrap.component.html',
})
export class BootstrapComponent {
  title = 'Talim';

  constructor(
  ) {
    console.log('test');
    this.getUser()
  }

  async getUser(): Promise<void> {
    const rsp = await fetch('/api/index?output=pb', {
      method: 'GET',
    })
    const ab = await rsp.arrayBuffer();
    const ua = new Uint8Array(ab);

    protobuf.util.Long = Long;
    protobuf.configure();

    const a = pb.UserList.decode(ua);

    for (const v of a.list.splice(1000, 10)) {
      let uid = v?.user?.uid;
      if (!uid) {
        continue;
      }
      uid = uid as Long;
      console.log(uid.toString());
    }
  }
}
