export namespace backend {
	
	export class SelectableTool {
	    toolName: string;
	    wslPath: string;
	    isSelected: boolean;
	
	    static createFrom(source: any = {}) {
	        return new SelectableTool(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.toolName = source["toolName"];
	        this.wslPath = source["wslPath"];
	        this.isSelected = source["isSelected"];
	    }
	}
	export class WslDistro {
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new WslDistro(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	    }
	}

}

