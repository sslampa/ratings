import chai from 'chai';
import React from 'react';
import {App} from '../../client/components/App';
import {List} from '../../client/components/List';
import {shallow, mount} from  'enzyme';

let expect = chai.expect;

describe('<App />',  () => {
  it('renders', () => {
    const component = shallow(<App />);
    expect(component.exists()).equal(true);
  });

  it('contains <List /> component', () => {
    const component = mount(<App />);
    expect(component.find(List)).to.have.length(1);
  })
})
