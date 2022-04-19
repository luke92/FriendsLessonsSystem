import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { User } from '../interfaces/user';
import { ConfigService } from './config.service';

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

  baseUrl: string = '';

  constructor(private http: HttpClient, private _configService: ConfigService) {
      this.baseUrl = this._configService.settings.apiUrl;
   }

  getUsers(): Observable<any> {
    const url = this.baseUrl + '/api/users';
    //Devuelve una copia del array MOCK
    //return this.listUsers.slice();
    //API
    return this.http.get(url);
  }

  getFriendships(): Observable<any> {
    const url = this.baseUrl + '/api/friendships';
    //Devuelve una copia del array MOCK
    //return this.listFriendships.slice();
    //API
    return this.http.get(url);
  }

  getLessonsByUserId(id: string): Observable<any> {
    const url = this.baseUrl + '/api/users/' + id + '/lessons';
    //Devuelve una copia del array MOCK
    //return this.listFriendships.slice();
    //API
    return this.http.get(url);
  }

  getFriendsByUserId(id: string): Observable<any> {
    const url = this.baseUrl + '/api/users/' + id + '/friends';
    //Devuelve una copia del array MOCK
    //return this.listFriendships.slice();
    //API
    return this.http.get(url);
  }
}
