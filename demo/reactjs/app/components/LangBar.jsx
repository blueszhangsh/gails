import React from 'react'
import {Nav, NavItem} from 'react-bootstrap';

const Widget = React.createClass({
    render: function() {
        return (
            <Nav pullRight>
                <NavItem href="/?locale=en-US" target="_blank">
                    English
                </NavItem>
                <NavItem href="/?locale=zh-CN" target="_blank">
                    简体中文
                </NavItem>
            </Nav>
        );
    }
});

export default Widget;
