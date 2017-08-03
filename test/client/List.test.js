import chai from 'chai';
import React from 'react';
import {List} from '../../client/components/List'
import {shallow} from  'enzyme';

let expect = chai.expect;

describe('<List />',  () => {
  it('renders', () => {
    const component = shallow(<List />);
    expect(component.exists()).equal(true);
  });
})
