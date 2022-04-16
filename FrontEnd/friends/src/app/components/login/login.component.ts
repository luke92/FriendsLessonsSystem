import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  form: FormGroup;
  loading = false;

  constructor(private fb: FormBuilder, private _snackBar: MatSnackBar, private router: Router) {
    this.form = this.fb.group({
      user: ['', Validators.required],
      password: ['', Validators.required],
    })
   }

  ngOnInit(): void {
  }

  login() {
    const user = this.form.value.user;
    const password = this.form.value.password; 
    
    if (user == 'admin' && password == 'admin123'){
      this.loginOk();
    }
    else {
      this.loginError();
    }
  }

  loginOk(){
    this.fakeLoading();
  }

  loginError() {
    this._snackBar.open('Invalid credentials', '', {
      horizontalPosition: 'center',
      verticalPosition: 'bottom',
      duration: 3000
    })
  }

  fakeLoading() {
    this.loading = true;
    setTimeout(() =>  {
      this.router.navigate(['dashboard']);
    }, 1500);
  }

}
