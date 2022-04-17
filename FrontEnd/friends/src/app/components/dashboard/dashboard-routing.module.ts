import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { DashboardComponent } from './dashboard.component';
import { FriendshipsComponent } from './friendships/friendships.component';
import { InitComponent } from './init/init.component';
import { UserDetailsComponent } from './users/user-details/user-details.component';
import { UsersComponent } from './users/users.component';

const routes: Routes = [
  { path: '', component: DashboardComponent, children: [
    { path: '', component: InitComponent},
    { path: 'users', component: UsersComponent},
    { path: 'friendships', component: FriendshipsComponent},
    { path: 'user-details/:id', component: UserDetailsComponent}
  ]}
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class DashboardRoutingModule { }
