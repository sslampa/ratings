import chai from 'chai';
import React from 'react';
import {App} from '../../client/components/App';
import {shallow, mount} from  'enzyme';

let expect = chai.expect;

describe('<App />',  () => {
  it('renders', () => {
    const component = shallow(<App />);
    expect(component.exists()).equal(true);
  });

  it('calls componentDidMount', () => {
    const component = mount(<App />);
  })
})
