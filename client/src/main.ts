import { bootstrapApplication } from '@angular/platform-browser';
import { appConfig } from './app/common/config';
import { BootstrapComponent } from './app/common/bootstrap.component';

bootstrapApplication(BootstrapComponent, appConfig)
  .catch((err) => console.error(err));
