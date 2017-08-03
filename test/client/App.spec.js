import chai from 'chai';
import React from 'react';
import {App} from '../../client/components/App'
import {shallow} from  'enzyme';

let expect = chai.expect;

describe('<App />',  () => {
  it('renders', () => {
    const component = shallow(<App />);
    expect(component.exists()).equal(true);
  });
})
