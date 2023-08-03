import { actionTypes } from "./actions"

const initialState = {
    User: {},
    Cart : [],
    filters : {
        discount: 0,
        gender : "male",
        price : {
            min : 100,
            max : 100
        }
    }
}
const {BUY_PRODUCT,ADD_CART,CHANGE_DISCOUNT_FILTER,
    CHANGE_GENDER_FILTER,CHANGE_PRICE_RANGE} = actionTypes;
export const Reducer = (state=initialState,action) =>{
    switch(action.type){
        case BUY_PRODUCT:
            return {
                ...state,
            }
        case ADD_CART:
            return {
                ...state,
                Cart: state.Cart.push(action.payload)
            }
        case CHANGE_DISCOUNT_FILTER :
            return {
                ...state,
                filters : {
                    ...state.filters,
                    discount: action.payload.discount
                }
            }


        case CHANGE_GENDER_FILTER : 
            return {
                ...state,
                filters : {
                    ...state.filters,
                    gender : action.payload.gender
                }
            }

        case CHANGE_PRICE_RANGE :
            return {
                ...state,
                filters : {
                    ...state.filters,
                    price : {
                     min : action.payload.min,
                     max : action.payload.max   
                    }
                }
            }


        default:
            return state
    }
}