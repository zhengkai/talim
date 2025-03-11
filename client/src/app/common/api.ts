import { pb } from '../../pb';
import Long from 'long';

class Api {

	userList: pb.UserRow[] = [];

	error451 = false;

	async fetch(uri: string): Promise<Uint8Array | null> {

		if (uri.includes('?')) {
			uri += `&`;
		} else {
			uri += `?`;
		}
		uri += 'output=pb';

		const rsp = await fetch(`/api/${uri}`, {
			method: 'GET',
		})
		if (!rsp?.ok) {
			if (rsp?.status === 451) {
				this.error451 = true;
			}
			return null;
		}
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
