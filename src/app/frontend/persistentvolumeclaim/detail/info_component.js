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

/**
 * @final
 */
<<<<<<< HEAD
export default class PersistentVolumeClaimInfoController {
=======
class PersistentVolumeClaimInfoController {
>>>>>>> upstream/master
  /**
   * Constructs persistent volume claim controller info object.
   */
  constructor() {
    /**
     * Persistent Volume Claim details. Initialized from the scope.
     * @export {!backendApi.PersistentVolumeClaimDetail}
     */
    this.persistentVolumeClaim;
  }
}

/**
 * Definition object for the component that displays persistent volume claim info.
 *
 * @return {!angular.Component}
 */
export const persistentVolumeClaimInfoComponent = {
  controller: PersistentVolumeClaimInfoController,
  templateUrl: 'persistentvolumeclaim/detail/info.html',
  bindings: {
    /** {!backendApi.PersistentVolumeClaimDetail} */
    'persistentVolumeClaim': '=',
  },
};
