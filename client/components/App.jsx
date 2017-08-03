import React from 'react';
import {List} from './List.jsx';

export class App extends React.Component {
  render() {
    return (
      <div>
        <h1>Hello World and Stephen</h1>
        <List items={[{id: 1, title: 'Louie', year: '2009', desc: 'A comedy show'}]} />
      </div>
    );
  }
}
