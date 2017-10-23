<<<<<<< HEAD
// Copyright 2017 The Kubernetes Dashboard Authors.
=======
// Copyright 2017 The Kubernetes Authors.
>>>>>>> upstream/master
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import {Validator} from '../../common/validators/validator';

/**
 * @final
 * @extends {Validator}
 */
export class LabelKeyNamePatternValidator extends Validator {
  constructor() {
    super();
    /** @type {!RegExp} */
    this.labelKeyNamePattern = /^([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9]$/;
  }

  /**
   * Returns true if the label key name (after the "/" if there is one) matches an alphanumeric
   * character (upper or lower case) optionally followed by alphanumeric or -_. and ending
   * with an alphanumeric character (upper or lower case), otherwise returns false.
   *
   * @override
   */
  isValid(value) {
    /** @type {number} */
    let slashPosition = value.indexOf('/');
    /** @type {string} */
    let labelKeyName = slashPosition > -1 ? value.substring(slashPosition + 1) : value;

    return this.labelKeyNamePattern.test(labelKeyName) || value === '';
  }
}
