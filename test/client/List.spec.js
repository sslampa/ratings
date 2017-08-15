import chai from 'chai';
import React from 'react';
import {List} from '../../client/components/List'
import {shallow} from  'enzyme';
let expect = chai.expect;

const items = [
  {
    id: 1,
    title: 'Louie',
    year: '2009',
    desc: 'A comedy show'
  },
  {
    id: 2,
    title: 'Breaking Bad',
    year: '2009',
    desc: 'A dramatic show'
  }
]

describe('<List />',  () => {
  it('renders', () => {
    const component = shallow(<List items={items}/>);
    expect(component.exists()).equal(true);
  });

  it('renders one item', () => {
    const component = shallow(<List items={[items[0]]}/>);
    expect(component.find('.list-item')).to.have.length(1);
  });

  it('renders mulitple items', () => {
    const component = shallow(<List items={items}/>);
    expect(component.find('.list-item')).to.have.length(2);
  })
})
