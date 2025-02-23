import { pb } from '../../pb';
import Long from 'long';

class Api {

  uuid: string = '54c2d184-77eb-4333-abfa-91b002e76827';

  userList: pb.UserRow[] = [];

  setUUID(uuid: string) {
    this.uuid = uuid;
  }

  async fetch(uri: string): Promise<Uint8Array> {

    if (uri.includes('?')) {
      uri += `&`;
    } else {
      uri += `?`;
    }
    uri += this.uuid;

    const rsp = await fetch(`/api/${uri}&output=pb`, {
      method: 'GET',
    })
    const ab = await rsp.arrayBuffer();
    return new Uint8Array(ab);
  }

  async index(): Promise<pb.UserRow[]> {
    const ua = await this.fetch('index');
    if (!ua?.length) {
      return [];
    }
    const o = pb.UserList.decode(ua)
    const li = this.userList;
    li.length = 0;
    for (const u of o.list) {
      if (!u) {
        continue;
      }
      li.push(pb.UserRow.fromObject(u));
    }
    return li;
  }

  async recent(): Promise<pb.TweetList | null> {
    const ua = await this.fetch('recent');
    if (!ua?.length) {
      return null;
    }
    return pb.TweetList.decode(ua);
  }

  async tweet(uid: Long, tid: Long): Promise<pb.TweetList | null> {
    const ua = await this.fetch(`tweet?uid=${uid}&tid=${tid}`);
    if (!ua?.length) {
      return null;
    }
    return pb.TweetList.decode(ua);
  }
}

export const api = new Api();
