import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { DashboardRoutingModule } from './dashboard-routing.module';
import { SharedModule } from '../shared/shared.module';
import { DashboardComponent } from './dashboard.component';
import { InitComponent } from './init/init.component';
import { NavbarComponent } from './navbar/navbar.component';
import { UsersComponent } from './users/users.component';
import { FriendshipsComponent } from './friendships/friendships.component';
import { UserDetailsComponent } from './users/user-details/user-details.component';


@NgModule({
  declarations: [
    DashboardComponent,
    InitComponent,
    NavbarComponent,
    UsersComponent,
    FriendshipsComponent,
    UserDetailsComponent
  ],
  imports: [
    CommonModule,
    DashboardRoutingModule,
    SharedModule
  ]
})
export class DashboardModule { }
