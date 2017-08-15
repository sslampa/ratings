import React from 'react';
import {List} from './List.jsx';

export class App extends React.Component {
  render() {
    return (
      <div>
        <h1>Hello World</h1>
        <List items={data} />
      </div>
    );
  }
}

const data = [
  {
    "id": 1,
    "title": "Louie",
    "year": "2009",
    "desc": "A comedy show",
    "rating": "5",
  },
  {
    "id": 2,
    "title": "Breaking Bad",
    "year": "2009",
    "desc": "A dramatic show",
    "rating": "5",
  },
  {
    "id": 3,
    "title": "Louie",
    "year": "2009",
    "desc": "A comedy show",
    "rating": "5",
  },
  {
    "id": 4,
    "title": "Louie",
    "year": "2009",
    "desc": "A comedy show",
    "rating": "5",
  },
  {
    "id": 5,
    "title": "Louie",
    "year": "2009",
    "desc": "A comedy show",
    "rating": "5",
  },
  {
    "id": 6,
    "title": "Louie",
    "year": "2009",
    "desc": "A comedy show",
    "rating": "5",
  },
  {
    "id": 7,
    "title": "Louie",
    "year": "2009",
    "desc": "A comedy show",
    "rating": "5",
  },
]
