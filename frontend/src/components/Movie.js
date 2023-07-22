// this is to show one movie displaying
import "../App.css"
const Movie = (props) =>{
    let x = props.movie;
     return (
        <div className="movie">
          
            <img src={x.url}></img>
            <h3>{x.name}</h3>
            <h5>{x.director}</h5>
        </div>
     )
}
export default Movie;