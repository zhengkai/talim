import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import Long from 'long';
import { api } from '../common/api';
import protobuf from 'protobufjs';
import { RouterLink } from '@angular/router';

@Component({
	selector: 'app-root',
	imports: [
		RouterOutlet,
		RouterLink,
	],
	templateUrl: './bootstrap.component.html',
})
export class BootstrapComponent {

	title = 'Talim';

	loadDone = false;
	loadError = false;
	error451 = false;

	constructor(
	) {
		protobuf.util.Long = Long;
		protobuf.configure();
		this.getUser()
	}

	async getUser(): Promise<void> {
		const li = await api.index();
		this.error451 = api.error451;
		this.loadDone = true;
		if (!li.length) {
			this.loadError = true;
		}
	}
}
