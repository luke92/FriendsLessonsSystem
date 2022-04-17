import { Component, OnInit } from '@angular/core';
import { UserService } from 'src/app/services/user.service';

@Component({
  selector: 'app-friendships',
  templateUrl: './friendships.component.html',
  styleUrls: ['./friendships.component.css']
})
export class FriendshipsComponent implements OnInit {

  friendships: string[] = [];

  constructor(private _userService: UserService) { }

  ngOnInit(): void {
    this.loadFriendships();
  }

  loadFriendships() {
    this._userService.getFriendships().subscribe(data => {
      this.friendships = data;
    });
  }

}
