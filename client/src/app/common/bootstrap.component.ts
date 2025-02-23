import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet],
  templateUrl: './bootstrap.component.html',
  styles: [],
})
export class AppComponent {
  title = 'Talim';
}
