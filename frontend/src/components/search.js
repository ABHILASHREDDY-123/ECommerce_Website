import axios from "axios"
const SearchByName=(name,setMovies)=>{
    const link=`http://www.omdbapi.com/?s=${name}&apikey=654293a2`
    axios.get(link).then((res)=>{
        if(res.data.Search){
            console.log(res);
            setMovies(res.data.Search);
        }
    })
}
const Search = (props) => {
  return (
    <div>
      <input type="text" placeholder="Search.." onChange={(e)=>{SearchByName(e.target.value,props.setMovies)}}/>
    </div>
  );
};
export default Search;