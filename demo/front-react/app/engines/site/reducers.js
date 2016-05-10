import {  REFRESH } from '../constants'

const initState = {links:[]};

export function siteInfo(state = initState, action) {    
    switch (action.type) {
        case REFRESH:
            return action.info;
        default:
            return state;
    }
}
