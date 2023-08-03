export const actionTypes = {
    BUY_PRODUCT:'BUY_PRODUCT',
    ADD_CART : 'ADD_CART',
    CHANGE_GENDER_FILTER : 'CHANGE_GENDER_FILTER',
    CHANGE_DISCOUNT_FILTER : 'CHANGE_DISCOUNT_FILTER',
    CHANGE_PRICE_RANGE : 'CHANGE_PRICE_RANGE'
}

export const buyProduct = () =>{

    return {
        type : actionTypes.BUY_PRODUCT,
    }
}

export const addCart = () =>{
    return {
        type : actionTypes.ADD_CART,
    }
}

export const changeGenderFilter = (props) =>{
    return {
        type : actionTypes.CHANGE_GENDER_FILTER,
        payload: {
           gender  : props.gender
        }
    }
}

export const changePriceRange = (props) =>{
    return {
        type : actionTypes.CHANGE_PRICE_RANGE,
        payload: {
            min : props.min,
            max : props.max
    }
     }
}
export const changeDiscountFilter = (props) =>{
    return {
        type : actionTypes.CHANGE_DISCOUNT_FILTER,
        payload : {
            discount : props.discount
        }
    }
}
