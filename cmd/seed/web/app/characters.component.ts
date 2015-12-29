import {Component} from 'angular2/core';
import {Router} from 'angular2/router';
import {CharacterService} from './character.service';
import {Character} from './character';

@Component({
  selector: 'my-characters',
  template: `
    <h2>Select a Character</h2>
    <ul class="characters">
      <li *ngFor="#character of characters" (click)="onSelect(character)">
        <span class="badge">{{character.id}}</span> {{character.name}}
      </li>
    </ul>
    <h2 *ngIf="currentCharacter">
      {{currentCharacter.name | uppercase}} is my character
    </h2>
  `,
  styles: [`
    .characters {list-style-type: none; margin-left: 1em; padding: 0; width: 14em;}
    .characters li { cursor: pointer; }
    .characters li:hover {color: #369; background-color: #EEE; }
  `]
})
export class CharactersComponent {
  private _characters: Character[];
  public currentCharacter: Character;

  constructor(private _characterService: CharacterService) { }

  get characters() {
    return this._characters || this.getCharacters()
  }

  onSelect(character: Character) { this.currentCharacter = character; }

  /////////////////

  private getCharacters() {
    this._characters = [];

    this._characterService.getCharacters()
      .then(characters => this._characters = characters);

    return this._characters;
  }
}
