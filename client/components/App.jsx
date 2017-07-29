import React from 'react';

export default class App extends React.Component {
  constructor(props) {
    super(props);
    this.state = {count: 0};
    this.add = this.add.bind(this);
  }

  add() {
    this.setState({count: this.state.count += 1});
  }
  render() {
    return (
      <div>
        <h1>Hello World and Stephen</h1>
        <p>Count is {this.state.count}</p>
        <button onClick={this.add}>Click me</button>
      </div>
    );
  }
}
