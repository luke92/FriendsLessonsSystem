import { Component, OnInit, ViewChild } from '@angular/core';
import { MatPaginator } from '@angular/material/paginator';
import { MatSort } from '@angular/material/sort';
import { MatTableDataSource } from '@angular/material/table';
import { ActivatedRoute } from '@angular/router';
import { Lesson } from 'src/app/interfaces/lesson';
import { UserService } from 'src/app/services/user.service';

@Component({
  selector: 'app-user-details',
  templateUrl: './user-details.component.html',
  styleUrls: ['./user-details.component.css']
})
export class UserDetailsComponent implements OnInit {

  id : string = '';
  friends: string[] = [];
  listLessons: Lesson[] = [];
  private retrivedata: Array<any> = [];

  displayedColumns: string[] = ['id', 'name', 'type'];
  //Cuando tiene el ! es un operador de not null assertion
  dataSource = new MatTableDataSource<any>();

  @ViewChild(MatPaginator) paginator!: MatPaginator;
  @ViewChild(MatSort) sort!: MatSort;

  constructor(private router: ActivatedRoute, private _userService: UserService) { }

  ngOnInit(): void {
    this.router.params.subscribe( p => {
      this.id = p['id'];
    });
    this.loadFriends(this.id);
    this.loadLessons(this.id);
  }

  loadFriends(id: string) {
    this._userService.getFriendsByUserId(id).subscribe(data => {
      this.friends = data;
    });
  }

  loadLessons(id: string) {      
    this._userService.getLessonsByUserId(id).subscribe(data => {
      this.retrivedata = data;
      for(let entry of this.retrivedata) {
        let lesson = {
          id: entry.id,
          name: entry.name,
          type: entry.type
        };
        this.listLessons.push(lesson);
      }
    });
    this.dataSource = new MatTableDataSource(this.listLessons);
  }

  ngAfterViewInit() {
    this.dataSource.paginator = this.paginator;
    this.dataSource.sort = this.sort;
  }

  applyFilter(event: Event) {
    const filterValue = (event.target as HTMLInputElement).value;
    this.dataSource.filter = filterValue.trim().toLowerCase();
  }
}
