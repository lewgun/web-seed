import {bootstrap} from 'angular2/platform/browser';
import {ROUTER_PROVIDERS} from 'angular2/router';
import {CharacterService} from './character.service';
import {AppComponent} from './app.component';

bootstrap(AppComponent, [
	ROUTER_PROVIDERS,
	CharacterService,
]);
