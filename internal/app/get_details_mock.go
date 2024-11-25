package app

type songDetail struct {
	releaseDate string `json:"releaseDate"`
	text        string `json:"text"`
	link        string `json:"link"`
}

func getSongDetail(group, song string) songDetail {

	return songDetail{
		releaseDate: "19700101",
		text:        "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas ut arcu magna. Aliquam sodales lacus non ipsum fermentum, id porttitor turpis pellentesque. Nulla at ipsum commodo, cursus mi non, tincidunt nibh. Praesent at nulla in neque ultrices congue in et sapien. Vestibulum id egestas lacus. Fusce porttitor aliquet enim non egestas. Suspendisse condimentum sit amet massa pellentesque efficitur. Donec felis mi, placerat non pretium vel, finibus tincidunt ipsum. ",
		link:        "localhost",
	}
}
