import {Component} from 'angular2/core';

@Component({
  selector: 'my-dashboard',
	template: `
		<h2>Dashboard</h2>
		<div>Hello {{name}}</div>
		<input [(ngModel)]="name" />
		<button (click)="sayHello()">Say Hello</button>
		<p>{{message}}</p>
	`
})
export class DashboardComponent {
	public name = 'john';
	public message = '';

	sayHello() {
		this.message = 'Hello ' + this.name + '!';
	}
}
