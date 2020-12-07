const searchBar = document.getElementById('searchBar');

const disFilterName = document.getElementById('disFilterName');

searchBar.addEventListener('keyup', (e) => {
    const searchString = e.target.value.toLowerCase();
    if (searchString == "") {
        // remove list under serach bar
        disFilterName.innerHTML = '';
        return
    }
    const filteredCharacters = Artists.filter((character) => {
        return (
            character.name.toLowerCase().includes(searchString) ||
            character.firstAlbum.includes(searchString) ||
            character.creationDate.toString().includes(searchString) ||
            character.Concerts.toString().toLocaleLowerCase().includes(searchString) ||
            character.members.toString().toLocaleLowerCase().includes(searchString)
        );
    });
    const searchFilterName = Artists
        .filter((character) => {
            return (
                character.name.toLowerCase().includes(searchString)
            );
        })
        .map((character) => {
            return `
                <li class="filterLi">
                    <p>${character.name} &nbsp&nbsp-&nbsp&nbsp :Band</p>
                </li>
            `;
        })
        .join('');

    const searchFilterAlbum = Artists
        .filter((character) => {
            return (
                character.firstAlbum.includes(searchString)
            );
        })
        .map((character) => {
            return `
                <li class="filterLi">
                    <p>${character.firstAlbum} &nbsp&nbsp-&nbsp&nbsp ${character.name}&nbsp&nbsp:Album</p>
                </li>
            `;
        })
        .join('');

    const searchFilterCreation = Artists
        .filter((character) => {
            return (
                character.creationDate.toString().includes(searchString)
            );
        })
        .map((character) => {
            return `
            <li class="filterLi">
                <p>${character.creationDate} &nbsp&nbsp-&nbsp&nbsp ${character.name}&nbsp&nbsp:CreationDate</p>
            </li>
        `;
        })
        .join('');


    let arrOfMembers = [];
    Artists
        .filter((character) => {
            for (let i = 0; i < character.members.length; i++) {
                if (character.members[i].toLocaleLowerCase().includes(searchString)) {
                    return character.members
                }
            }
        })
        .forEach((ArtistsMembers) => {
            ArtistsMembers.members.filter((character) => {
                if (character.toLowerCase().includes(searchString)) {
                    arrOfMembers.push(character)
                    return character.toLowerCase().includes(searchString)
                }
            });
        })

    // display hint group name/album/creation date under search bar
    disFilterName.innerHTML = searchFilterName + searchFilterAlbum + searchFilterCreation + displayFilterMembers(arrOfMembers);

    displayCharacters(filteredCharacters);
});

// display hint Group members name
const displayFilterMembers = (Artists) => {
    const htmlString = Artists
        .map((character) => {
            return `
            <li class="filterLi">
                <p>${character.toString()} &nbsp&nbsp-&nbsp&nbsp :Members</p>
            </li>
        `;
        })
        .join('');
    return htmlString;
};

const ul = document.getElementById('wrapperElements');

// display filtered characters
const displayCharacters = (Artists) => {
    const htmlString = Artists
        .map((character) => {
            return `
            <li class="scene">
                <div class="movie" onclick="return true">
                    <div class="poster">
                        <img src="${character.image}" height="260px" width="260px"></img>
                        <h1>${character.name}</h1>
                    </div>
                    <div class="info">
                        <header>
	    					<span class="year">Creation Date: &nbsp </span>
		    				<span class="rating">${character.creationDate}<br></span><br>
    		    			<span class="duration">First Album: &nbsp</span>
		    	    		<span class="rating">${character.firstAlbum}</span> <hr>
                            <span class="duration">Members:</span><br>                      
						    <span class="members">${character.members}</span><br>							   
					    </header>
					    <p>Concerts: <br> 
						    <span class="concerts">${character.Concerts}</span><br>
					    </p>
                    </div>
                </div>
            </li>
      `;
        })
        .join('');
    ul.innerHTML = htmlString;
};

displayCharacters(Artists);