import { Component } from '@angular/core';

@Component({
    selector: 'my-app',
    template: '<h1>{{title}}</h1><h2>{{note.content}}详细信息</h2>'
})

export class AppComponent {
  title = "测试";
  note: Note = {
    id:1,
    content:'笔记1'
  };
}

export class Note{
  id: number;
  content: string;
}
