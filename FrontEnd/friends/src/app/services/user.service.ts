import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { User } from '../interfaces/user';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  listUsers: User[] = [
    {id: "1", name: 'Joe', lastname: 'Wilson', username: "joe", email: "joe@gmail.com"},
    {id: "2", name: 'Joe', lastname: 'Miller', username: "joe", email: "joe@gmail.com"},
  ];

  constructor(private http: HttpClient) { }

  getUsers(): Observable<any> {
    const url = '/api/users';
    //Devuelve una copia del array MOCK
    //return this.listUsers.slice();
    //API
    return this.http.get(url);
  }
}
