import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Config } from '../interfaces/config';

@Injectable({
  providedIn: 'root'
})
export class ConfigService {
  
  private configSettings: any = null;

  constructor(private http: HttpClient) { }

  get settings() {
    return this.configSettings;
  }
  
  public loadConfig(): Promise<any> {
    return new Promise((resolve, reject) => {
      this.http.get<Config>('./assets/config.json').subscribe((response: any) => {
        this.configSettings = response;
        resolve(true);
      });
    });
  }
}
