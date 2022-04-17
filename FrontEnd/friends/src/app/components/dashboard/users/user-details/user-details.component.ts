import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { UserService } from 'src/app/services/user.service';

@Component({
  selector: 'app-user-details',
  templateUrl: './user-details.component.html',
  styleUrls: ['./user-details.component.css']
})
export class UserDetailsComponent implements OnInit {

  id : string = '';

  constructor(private router: ActivatedRoute, private _userService: UserService) { }

  ngOnInit(): void {
    this.router.params.subscribe( p => {
      this.id = p['id'];
    });
  }

}
