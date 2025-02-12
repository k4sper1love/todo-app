export namespace main {
	
	export class Task {
	    id: number;
	    text: string;
	    status: string;
	
	    static createFrom(source: any = {}) {
	        return new Task(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.text = source["text"];
	        this.status = source["status"];
	    }
	}

}

