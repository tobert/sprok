// +build darwin
// OSX and some others do not offer setres[gu]id

package sprok

/*
 * Copyright 2015 Albert P. Tobey <atobey@datastax.com> @AlTobey
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import (
	"log"
	"syscall"
)

func (p *Process) drop_privileges() {
	var err error

	if p.Gid >= 0 {
		err = syscall.Setregid(p.Gid, p.Gid)
		if err != nil {
			log.Fatalf("setregid(%d...) failed: %s\n", p.Gid, err)
		}
	}

	if p.Uid >= 0 {
		err = syscall.Setreuid(p.Uid, p.Uid)
		if err != nil {
			log.Fatalf("setreuid(%d...) failed: %s\n", p.Uid, err)
		}
	}
}

