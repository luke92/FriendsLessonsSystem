import { Component, OnInit, ViewChild } from '@angular/core';
import { MatPaginator } from '@angular/material/paginator';
import { MatSort } from '@angular/material/sort';
import { MatTableDataSource } from '@angular/material/table';
import { Router } from '@angular/router';
import { User } from 'src/app/interfaces/user';
import { UserService } from 'src/app/services/user.service';
@Component({
  selector: 'app-users',
  templateUrl: './users.component.html',
  styleUrls: ['./users.component.css']
})
export class UsersComponent implements OnInit {
  
  listUsers: User[] = [];
  private retrivedata: Array<any> = [];

  displayedColumns: string[] = ['id', 'name', 'lastname', 'username', 'email', 'actions'];
  //Cuando tiene el ! es un operador de not null assertion
  dataSource = new MatTableDataSource<any>();

  @ViewChild(MatPaginator) paginator!: MatPaginator;
  @ViewChild(MatSort) sort!: MatSort;

  constructor(private _userService: UserService, private router: Router) { }

  ngOnInit(): void {
    this.loadUsers();
  }

  loadUsers() {
    this._userService.getUsers().subscribe(data => {
      this.retrivedata = data;
      for(let entry of this.retrivedata) {
        let user = {
          id: entry.id,
          name: entry.name,
          lastname: entry.lastname,
          username: entry.username,
          email: entry.email
        };
        this.listUsers.push(user);
      }
    });
    this.dataSource = new MatTableDataSource(this.listUsers);
  }

  ngAfterViewInit() {
    this.dataSource.paginator = this.paginator;
    this.dataSource.sort = this.sort;
  }

  applyFilter(event: Event) {
    const filterValue = (event.target as HTMLInputElement).value;
    this.dataSource.filter = filterValue.trim().toLowerCase();
  }

  userDetails(id: string) {
    console.log(id);
    this.router.navigate(['/dashboard/user-details', id]);
  }

}
