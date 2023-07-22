import Movie from "./Movie"
import "../App.css";
const DisplayMovies = (props) =>{
      let movies = props.movies;
      console.log(movies);
      
      return (
        <div className="movies">
           {
            movies.map(x => {
              return <Movie movie={x}/>
            })
           }
        </div>

      )


}
export default DisplayMovies