import { Injectable } from '@angular/core';
import { User } from '../interfaces/user';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  listUsers: User[] = [
    {id: "1", fullname: 'Joe', username: "joe", email: "joe@gmail.com"},
    {id: "2", fullname: 'Joe', username: "joe", email: "joe@gmail.com"},
  ];

  constructor() { }

  getUsers() {
    //Devuelve una copia del array
    return this.listUsers.slice();
  }
}
