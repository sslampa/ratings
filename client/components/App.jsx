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
    "year": "2008",
    "desc": "A dramatic show",
    "rating": "5",
  },
  {
    "id": 3,
    "title": "Rick and Morty",
    "year": "2013",
    "desc": "A comedy show",
    "rating": "5",
  },
  {
    "id": 4,
    "title": "Frasier",
    "year": "1993",
    "desc": "A comedy show",
    "rating": "5",
  },
  {
    "id": 5,
    "title": "Community",
    "year": "2009",
    "desc": "A comedy show",
    "rating": "5",
  },
  {
    "id": 6,
    "title": "Mad Men",
    "year": "2007",
    "desc": "A dramatic show",
    "rating": "5",
  },
  {
    "id": 7,
    "title": "30 Rock",
    "year": "2006",
    "desc": "A comedy show",
    "rating": "5",
  },
]
