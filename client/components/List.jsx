import React from 'react';

export class List extends React.Component {
  makeList = () => {
    let list = this.props.items.map(item => {
      return this.itemPanel(item)
    })

    return list
  }

  itemPanel = (item) => {
    return (
      <div className='list-item' key={item.id}>
        <div className='list-item-header'>
          <h3>{item.title} - {item.year}</h3>
        </div>
        <div className='list-item-body'>
          <p>{item.desc}</p>
        </div>
      </div>
    );
  }

  render() {
    return (
      <div>
        {this.makeList()}
      </div>
    )
  }
}
