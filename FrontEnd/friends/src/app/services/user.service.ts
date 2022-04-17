import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { User } from '../interfaces/user';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  //MOCKS
  listUsers: User[] = [
    {id: "1", name: 'Joe', lastname: 'Wilson', username: "joe", email: "joe@gmail.com"},
    {id: "2", name: 'Joe', lastname: 'Miller', username: "joe", email: "joe@gmail.com"},
  ];

  listFriendships: string[] = [
    "Joe",
    "Rachel"
  ];

  listLessons: string[] = [
    "Math",
    "Spanish"
  ];

  constructor(private http: HttpClient) { }

  getUsers(): Observable<any> {
    const url = '/api/users';
    //Devuelve una copia del array MOCK
    //return this.listUsers.slice();
    //API
    return this.http.get(url);
  }

  getFriendships(): Observable<any> {
    const url = '/api/friendships';
    //Devuelve una copia del array MOCK
    //return this.listFriendships.slice();
    //API
    return this.http.get(url);
  }

  getLessonsByUserId(id: string): Observable<any> {
    const url = '/api/users/' + id + '/lessons';
    //Devuelve una copia del array MOCK
    //return this.listFriendships.slice();
    //API
    return this.http.get(url);
  }

  getFriendsByUserId(id: string): Observable<any> {
    const url = '/api/users/' + id + '/friends';
    //Devuelve una copia del array MOCK
    //return this.listFriendships.slice();
    //API
    return this.http.get(url);
  }
}
