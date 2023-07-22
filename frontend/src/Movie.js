const Movie = (props)=>{
   let each_movie=props.movie
   return (
     <div className="movie">
        <h3>{each_movie.name}</h3>
        <h3>{each_movie.gender}</h3>
        <h3>{each_movie.age}</h3>
        <img src={each_movie.url} height="300px"/>
     </div>
   ) 
}
export default Movie