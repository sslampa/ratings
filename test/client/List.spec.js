import chai from 'chai';
import React from 'react';
import {List} from '../../client/components/List'
import {shallow} from  'enzyme';

let expect = chai.expect;
const items = [
  {
    title: 'Louie',
    year: '2009',
    desc: 'A comedy show'
  }
]

describe('<List />',  () => {
  it('renders', () => {
    const component = shallow(<List items={items}/>);
    expect(component.exists()).equal(true);
  });

  it('renders one item', () => {
    const component = shallow(<List items={items}/>);
    expect(component.find('.list-item')).to.have.length(1);
  });
})
