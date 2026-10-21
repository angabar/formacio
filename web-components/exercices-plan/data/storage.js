export default {
    get saved() {
        const savedPlans = [];

        Object.keys(localStorage).forEach((key) => {
            savedPlans.push(key);
        });

        return savedPlans;
    },
    save(name) {
        localStorage.setItem(name, JSON.stringify(this._currentWorkout));
    },
    load(key) {
        this._currentWorkout = JSON.parse(localStorage.getItem(key));
    },
    edit(id, key, value) {
        let exercise;

        for (let c = 0; c < this._currentWorkout.length; c++) {
            if (id === this._currentWorkout[c].id) {
                exercise = this._currentWorkout[c];
                exercise[key] = value;
            }
        }
    },
    add(exercise) {
        if (!this._currentWorkout) {
            this._currentWorkout = [];
        }

        exercise.id = this.createID();
        this._currentWorkout.push(exercise);
    },
    remove(id) {
        if (!this._currentWorkout) {
            return;
        }

        for (let c = 0; c < this._currentWorkout.length; c++) {
            if (id === this._currentWorkout[c].id) {
                const deleted = this._currentWorkout.splice(c, 1);
                return;
            }
        }
    },
    clear() {
        this._currentWorkout = [];
    },
    get exercises() {
        if (!this._currentWorkout) {
            this._currentWorkout = [];
        }

        return this._currentWorkout;
    },
    get totalDuration() {
        let ttlTime = 0;

        for (let c = 0; c < this._currentWorkout.length; c++) {
            ttlTime += this.getDurationOfExercise(this._currentWorkout[c]);
        }

        return ttlTime;
    },
    formatTime(seconds) {
        return new Date(1000 * seconds).toISOString().substring(11, 8);
    },
};
