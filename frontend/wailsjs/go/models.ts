export namespace main {
	
	export class BrewConfig {
	    coreUrl: string;
	    bottleUrl: string;
	
	    static createFrom(source: any = {}) {
	        return new BrewConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.coreUrl = source["coreUrl"];
	        this.bottleUrl = source["bottleUrl"];
	    }
	}
	export class OperationResult {
	    success: boolean;
	    message: string;
	    output: string;
	
	    static createFrom(source: any = {}) {
	        return new OperationResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.success = source["success"];
	        this.message = source["message"];
	        this.output = source["output"];
	    }
	}
	export class Package {
	    name: string;
	    version: string;
	    description: string;
	    type: string;
	    installed: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Package(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.version = source["version"];
	        this.description = source["description"];
	        this.type = source["type"];
	        this.installed = source["installed"];
	    }
	}

}

