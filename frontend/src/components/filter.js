import { changeDiscountFilter, changeGenderFilter, changePriceRange } from "../redux/actions";
import "../styles/filter.css"
import store from './../redux/store';
const GenderFilter = ({ gender }) => {

    return (
        <div className="genderfilter">
            <span>Gender : </span>
            <div className="checkboxa">
                <span className="checkboxspan">Male</span>
                <input className="inputcheck" type="checkbox" checked={(gender === "male") ? true : false} onClick={(e) => {
                    if (e.target.checked) {
                        const obj = {
                            gender: "male"
                        }
                        store.dispatch(changeGenderFilter(obj));
                    }
                    else {

                        const obj = {
                            gender: "none"
                        }
                        store.dispatch(changeGenderFilter(obj))

                    }
                }} />
            </div>

            <div className="checkboxa">
                <span className="checkboxspan">Female</span>
                <input className="inputcheck" type="checkbox" checked={(gender === "female") ? true : false} onClick={(e) => {
                    if (e.target.checked) {
                        const obj = {
                            gender: "female"
                        }
                        store.dispatch(changeGenderFilter(obj));
                    }
                    else {

                        const obj = {
                            gender: "gender"
                        }
                        store.dispatch(changeGenderFilter(obj))

                    }
                }} />
            </div>

        </div>
    )
}
const CostFilter = ({ mini, maxi }) => {
    return (
        <div className="costfilter">
            <span>Price : </span>
            <div className="selecta">
                <select id="selectvalue" value={mini} onChange={(e) => {
                    const obj = {
                        min: Number(e.target.value),
                        max: Math.max(Number(e.target.value), maxi)
                    }
                    store.dispatch(changePriceRange(obj));
                }}>
                    <option value={100}>100</option>
                    <option value={500}>500</option>
                    <option value={1000}>1000</option>
                    <option value={10000}>10000</option>
                </select>
            </div>
            -
            <div className="selecta">
                <select value={maxi} onChange={(e) => {

                    const obj = {
                        min: Math.min(mini, Number(e.target.value)),
                        max: Number(e.target.value)
                    }
                    store.dispatch(changePriceRange(obj));
                }}>

                    <option value={100}>100</option>
                    <option value={500}>500</option>
                    <option value={1000}>1000</option>
                    <option value={10000}>10000</option>

                </select>
            </div>

        </div>
    )
}
const updateValue = (value) => {
    store.dispatch(changeDiscountFilter({ discount: value }))
}
const DiscountFilter = ({ value }) => {
    return (
        <div className="discountfilter">
            <span >Discount : <span id="discount">{value} %</span> </span>
            <input type="range" min="0" max="100" step={1} defaultValue={value} style={{ width: "50%" }} onInput={(e) => { updateValue(e.target.value) }} />
        </div>
    )
}
const Filter = () => {
    const { filters } = store.getState();
    return (
        <div className="showFilter">
            <GenderFilter gender={filters.gender} />
            <CostFilter maxi={filters.price.max} mini={filters.price.min} />
            <DiscountFilter value={filters.discount} />
        </div>
    )
}
export default Filter;